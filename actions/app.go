package actions

import (
	"github.com/gobuffalo/buffalo"
	"github.com/gobuffalo/envy"
	"github.com/gobuffalo/mw-contenttype"
	"github.com/gobuffalo/mw-forcessl"
	"github.com/gobuffalo/mw-paramlogger"
	"github.com/unrolled/secure"

	"worldlocations/models"

	"github.com/gobuffalo/buffalo-pop/pop/popmw"
	"github.com/gobuffalo/x/sessions"
	"github.com/rs/cors"
)

// ENV is used to help switch settings based on where the
// application is being run. Default is "development".
var ENV = envy.Get("GO_ENV", "development")
var app *buffalo.App

// App is where all routes and middleware for buffalo
// should be defined. This is the nerve center of your
// application.
//
// Routing, middleware, groups, etc... are declared TOP -> DOWN.
// This means if you add a middleware to `app` *after* declaring a
// group, that group will NOT have that new middleware. The same
// is true of resource declarations as well.
//
// It also means that routes are checked in the order they are declared.
// `ServeFiles` is a CATCH-ALL route, so it should always be
// placed last in the route declarations, as it will prevent routes
// declared after it to never be called.
func App() *buffalo.App {
	if app == nil {
		app = buffalo.New(buffalo.Options{
			Env:          ENV,
			SessionStore: sessions.Null{},
			PreWares: []buffalo.PreWare{
				cors.Default().Handler,
			},
			SessionName: "_worldlocations_session",
		})

		// Automatically redirect to SSL
		app.Use(forceSSL())

		// Log request parameters (filters apply).
		app.Use(paramlogger.ParameterLogger)

		// Set the request content type to JSON
		app.Use(contenttype.Set("application/json"))

		// Wraps each request in a transaction.
		//  c.Value("tx").(*pop.Connection)
		// Remove to disable this.
		app.Use(popmw.Transaction(models.DB))

		//v1 groupe
		v1 := app.Group("/v1")
		// countries
		inquiry := Inquiry{}
		v1.GET("/countries/", v1Handler(inquiry.countryList))
		v1.GET("/countries/{alpha_2_code}/", v1Handler(inquiry.countryShow))
		v1.GET("/countries/{alpha_2_code}/names", v1Handler(inquiry.countryNames))
		v1.GET("/countries-names/", v1Handler(inquiry.countriesNamesList))
		//subdivision Names
		v1.GET("/subdivisions/names", v1Handler(inquiry.subdivisionNamesList))
		v1.GET("/subdivisions/{subdivision_code}/names/", v1Handler(inquiry.subdivisionNamesShow))
		//subdivision
		v1.GET("/subdivisions/", v1Handler(inquiry.subdivisionList))
		v1.GET("/subdivisions/{subdivision_code}/", v1Handler(inquiry.subdivisionShow))
		v1.GET("/countries/{alpha_2_code}/subdivisions/", v1Handler(inquiry.countrySubdivisions))
		//languages
		v1.GET("/languages/", v1Handler(inquiry.languagesList))
		v1.GET("/languages/{language_alpha_2_code}/", v1Handler(inquiry.languagesShow))
		// cities
		v1.GET("cities/", v1Handler(inquiry.citiesList))
		v1.GET("cities/{id}", v1Handler(inquiry.citiesShow))
		v1.GET("countries/{alpha_2_code}/cities/", v1Handler(inquiry.countryCities))
		v1.GET("subdivisions/{subdivision_code}/cities/", v1Handler(inquiry.subdivisionCities))

		//default home controller
		app.GET("/", HomeHandler)

	}

	return app
}

// forceSSL will return a middleware that will redirect an incoming request
// if it is not HTTPS. "http://example.com" => "https://example.com".
// This middleware does **not** enable SSL. for your application. To do that
// we recommend using a proxy: https://gobuffalo.io/en/docs/proxy
// for more information: https://github.com/unrolled/secure/
func forceSSL() buffalo.MiddlewareFunc {
	return forcessl.Middleware(secure.Options{
		SSLRedirect:     ENV == "production",
		SSLProxyHeaders: map[string]string{"X-Forwarded-Proto": "https"},
	})
}
