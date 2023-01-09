export default {
    title: 'CRM DOCS',
    description: 'CRM Docs',
    themeConfig: {
        siteTitle: 'CRM DOCS',
        nav: [
            { text: '文档', link: '/project/introduction' },
            { text: '更新日志', link: '/logs/update-log' },
            { text: '关于', link: '/about/about' },
            { text: '赞赏', link: '/sponsor/sponsor' }
        ],
        socialLinks: [
            { icon: 'github', link: 'https://github.com/zchengo/crm' }
        ],
        sidebar: {
            '/project/': [
                {
                    text: '文档',
                    collapsible: true,
                    items: [
                        { text: '简介', link: '/project/introduction' },
                        { text: '技术栈', link: '/project/technology-stack' },
                        { text: '快速运行', link: '/project/getting-started' },
                        { text: '部署指南', link: '/project/deploy-guide' },
                        { text: '许可证', link: '/project/open-license' }
                    ],
                },
            ]
        },
        footer: {
            copyright: 'Copyright © 2023-present zocrm.cloud'
        },
        lastUpdated: true,
        lastUpdatedText: 'Updated Date',
        // carbonAds: {
        //     code: 'your-carbon-code',
        //     placement: 'your-carbon-placement'
        // }
    }
}
