// configs depending on environment
let base = ''
let apiURL = "http://localhost:8080/query"

console.log(process.env.NODE_ENV)
if (process.env.NODE_ENV === "production") {
  apiURL = "/query"
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
    repo: 'advenjourney/advenjourney',
    docsBranch: 'main',
    docsDir: 'web/docs',
    editLinks: true,
    editLinkText: 'Edit this page',

    nav: [
      { 
        text: 'Create a Trip', 
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
      // '/trips/': [
      //   ''
      // ],
      '/about-us/': [
        '',
        'vision',
        'strategy',
        'community'
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