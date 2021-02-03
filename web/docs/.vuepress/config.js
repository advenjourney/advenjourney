module.exports = {

  lang: 'en-US',
  title: 'AdvenJourney',
  description: 'Just playing around.',

  themeConfig: {
    custom: {
      url: process.env.NODE_ENV === 'development' ? 'http://localhost:8081' : 'https:advenjourney.com/api'
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