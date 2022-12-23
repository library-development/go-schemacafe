package schemacafe

import "lib.dev/marketing"

var Marketing struct {
	HomePage *marketing.LandingPage
} = struct {
	HomePage *marketing.LandingPage
}{
	HomePage: &marketing.LandingPage{
		Title:   "schema.cafe",
		Heading: "A proper home for your data schemas",
	},
}
