export default {
    title: 'CrmDocs',
    description: 'CrmDocs',
    themeConfig: {
        logo: '/logo.svg',
        siteTitle: 'CrmDocs',
        algolia: {
            apiKey: 'your_api_key',
            indexName: 'index_name'
        },
        nav: [
            { text: '文档', link: '/project/docs/introduction' },
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
                        { text: '简介', link: '/project/docs/introduction' },
                        { text: '快速开始', link: '/project/docs/getting-started' },
                        { text: '部署指南', link: '/project/docs/deploy-guide' },
                        { text: '详细设计', link: '/project/docs/detailed-design' },
                        { text: '问题反馈', link: '/project/docs/problem-feedback' },
                        { text: '更新日志', link: '/project/docs/update-log' },
                    ]
                },
                {
                    text: '前端',
                    collapsible: true,
                    items: [
                        { text: '封装 axios 请求库', link: '/project/frontend/axios-package' },
                        { text: '页面中的加载进度条', link: '/project/frontend/nprogress' },
                    ]
                },
                {
                    text: '后端',
                    collapsible: true,
                    items: [
                        { text: '跨域请求处理', link: '/project/backend/cors-handle' },
                        { text: '用户授权与认证', link: '/project/backend/jwt-handle' },
                    ]
                },
                {
                    text: '运维',
                    collapsible: true,
                    items: [
                        { text: '部署到云服务器', link: '/project/devops/deploy-cloudserver' },
                        { text: '持续集成与交付 (CI/CD)', link: '/project/devops/ci-cd' },
                    ]
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
