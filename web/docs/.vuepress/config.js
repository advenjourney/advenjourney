// configs depending on environment
let base = ''
let apiURL = "http:localhost:9000/query"

console.log(process.env.NODE_ENV)
if (process.env.NODE_ENV === "production") {
  apiURL = "https://api.advenjourney.com/query"
}
if (process.env.API) {
  apiURL = process.env.API
}
if (process.env.BASE) {
  base = process.env.BASE
}
console.log("Base: ", base, "API: ", apiURL)

module.exports = {
  lang: 'en-US',
  title: 'AdvenJourney',
  description: 'Just playing around.',
  base: base,

  themeConfig: {
    custom: {
      url: apiURL,
    },
    repo: 'advenjourney/website',
    docsBranch: 'main',
    docsDir: 'docs',
    editLinks: true,
    editLinkText: 'Edit this page',

    nav: [
      { 
        text: 'Trips', 
        link: '/trips/', 
        activeMatch: '^/trips/' 
      },
      { 
        text: 'About Us', 
        link: '/about-us/', 
        activeMatch: '^/about-us/' 
      },
      {
        text: 'Login',
        link: '/login'
      }
    ],

    displayAllHeaders: true,
    sidebar: {
      '/trips/': getTripsSidebar(),
      '/about-us/': [
        '',
        'mission',
        'community',
        'team'
      ]
    }
  }
}

function getTripsSidebar() {
  return [
    '',
    'places',
    {
      title: 'Join a Trip',
      path: '/trips/join',
      collapsable: false,
    },
    {
      title: 'Create a Trip',
      collapsable: false,
      children: [
        'create',
        'publish'
      ]
    }
  ]
}