package analyze

type Analyze struct {
	userAgent  string
	apiVersion string
	apiPlugins string
	urlPlugins string

	indicators  []string
	hintPlugins []Plugin

	vIndicatorsRssFeed       VersionIndicator
	vIndicatorsMetaTag       VersionIndicator
	vIndicatorsEnquedScripts []VersionIndicator
	vIndicatorsLoginPage     []VersionIndicator

	data Store
}

type Store struct {
	htmlIndex string
}

type Plugin struct {
	name string
	uri  string
	key  string
	slug string
}

type VersionIndicator struct {
	indicator string
	split     string
}

func New() *Analyze {
	return &Analyze{
		userAgent:  "Mozilla/5.0 (compatible; GoWordlytics/1.0; +",
		apiVersion: "https://api.wordpress.org/core/version-check/1.7/",
		apiPlugins: "https://api.wordpress.org/plugins/info/1.0/%s.json",
		indicators: []string{
			"/wp-content",
			"/wp-includes",
			"/wp-admin",
			"/wp-json",
			"/mu-plugins",
			"name=\"generator\" content=\"WordPress",
			"wp-embed.min.js",
			"wp-emoji-release.min.js",
			"wp-emoji.min.js",
		},
		vIndicatorsMetaTag: VersionIndicator{
			indicator: "WordPress",
		},
		vIndicatorsRssFeed: VersionIndicator{
			indicator: "https://wordpress.org/?v=",
		},
		vIndicatorsEnquedScripts: []VersionIndicator{
			{
				indicator: "/wp-includes/js/wp-embed.min.js?ver=",
			},
			{
				indicator: "/wp-includes/css/dist/block-library/style.min.css?ver=",
			},
		},
		vIndicatorsLoginPage: []VersionIndicator{
			{
				indicator: "/wp-admin/css/forms.min.css?ver=",
			},
			{
				indicator: "/wp-admin/css/login.min.css?ver=",
			},
			{
				indicator: "/wp-admin/load-styles.php",
				split:     "ver=",
			},
		},
		hintPlugins: []Plugin{
			{
				name: "Yoast SEO",
				uri:  "https://de.wordpress.org/plugins/wordpress-seo/",
				key:  "Yoast SEO plugin",
				slug: "wordpress-seo",
			},
			{
				name: "Schema",
				uri:  "https://de.wordpress.org/plugins/schema/",
				key:  "schema.press",
				slug: "schema",
			},
			{
				name: "W3 Total Cache",
				uri:  "https://de.wordpress.org/plugins/w3-total-cache/",
				key:  "Performance optimized by W3 Total Cache",
				slug: "w3-total-cache",
			},
			{
				name: "WooCommerce",
				uri:  "https://de.wordpress.org/plugins/woocommerce/",
				key:  "woocommerce-no-js",
				slug: "woocommerce",
			},
			{
				name: "ShortPixel Image Optimizer",
				uri:  "https://de.wordpress.org/plugins/shortpixel-image-optimiser/",
				key:  "shortpixel-image-optimiser",
				slug: "shortpixel-image-optimiser",
			},
			{
				name: "KeyCDN",
				uri:  "https://www.keycdn.com/",
				key:  "KeyCDN",
				slug: "cdn-enabler",
			},
			{
				name: "Borlabs Cookie",
				uri:  "https://de.borlabs.io/borlabs-cookie/",
				key:  "borlabs-cookie",
				slug: "borlabs-cookie",
			},
			{
				name: "Gravityforms",
				uri:  "https://www.gravityforms.com/",
				key:  "gravityforms",
				slug: "gravityforms",
			},
			{
				name: "Wp Rocket",
				uri:  "https://wp-rocket.me/de/",
				key:  "wp-rocket",
				slug: "wp-rocket",
			},
			{
				name: "Autoptimize",
				uri:  "https://de.wordpress.org/plugins/autoptimize/",
				key:  "/cache/autoptimize/",
				slug: "autoptimize",
			},
			{
				name: "Cookie Notice & Compliance for GDPR / CCPA",
				uri:  "https://de.wordpress.org/plugins/cookie-notice/",
				key:  "<!-- / Cookie Notice plugin -->",
				slug: "cookie-notice",
			},
		},
	}
}
