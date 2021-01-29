module.exports = {
  lang: 'en-US',
  title: 'AdvenJourney',
  description: 'Just playing around.',

  themeConfig: {
    repo: 'advenjourney/site',
    docsDir: 'docs',

    editLinks: true,
    editLinkText: 'Edit this page',
    lastUpdated: 'Last Updated',

    // algolia: {
    //   apiKey: 'c57105e511faa5558547599f120ceeba',
    //   indexName: 'vitepress'
    // },

    // carbonAds: {
    //   carbon: 'CEBDT27Y',
    //   custom: 'CKYD62QM',
    //   placement: 'vuejsorg'
    // },

    nav: [
      {
        text: 'Home',
        link: '/',
        activeMatch: '^/$'
      },
      { 
        text: 'Trips', 
        link: '/trips/', 
        activeMatch: '^/trips/' 
      },
      { 
        text: 'About Us', 
        link: '/about-us/', 
        activeMatch: '^/about-us/' 
      }
    ],

    sidebar: {
      '/trips/': getTripsSidebar(),
      '/about-us/': getAboutUsSidebar(),
      '/': getAboutUsSidebar()
    }
  }
}

function getTripsSidebar() {
  return [
    'places',
    {
      title: 'Join a Trip',
      path: '/trips/join',
      collapsable: false,
      // initialOpenGroupIndex: 1,
      // sidebarDepth: 3
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

function getAboutUsSidebar() {
  return [
    {
      title: 'Mission',
      children: [
        '/about-us/'
      ]
    },
    'community',
    'team'
  ]
}
