// https://v2.vuepress.vuejs.org/reference/default-theme/config.html

import { defaultTheme } from '@vuepress/theme-default'
import { searchPlugin } from '@vuepress/plugin-search'
import { nprogressPlugin } from '@vuepress/plugin-nprogress'

let ruThemeConfig = {
    selectLanguageText: 'Языки',
    selectLanguageName: 'Русский',
    editLinkText: 'Редактировать',
    contributorsText: 'Авторы',
    lastUpdatedText: 'Обновлено',

    navbar: [
        { text: 'Быстрый старт', link: '/ru/guide/' },
        // { text: 'API', link: '/ru/api/' },
        // { text: 'Примеры', link: '/ru/examples/' },
        // { text: 'Ссылки', children: [
        //     { text: 'Extra', link: '/ru/extra/' },
        //     { text: 'Блог', link: '/ru/blog/' },
        //     { text: 'О нас', link: '/ru/about/' },
        // ]},

        // { text: 'ru', link: '/ru/guide/', activeMatch: '^/ru/' },
        // { text: 'en', link: '/en/guide/', activeMatch: '^/en/' }
    ],
};

let enThemeConfig = {
    selectLanguageText: 'Languages',
    selectLanguageName: 'English',

    navbar: [
        { text: 'Quick Start', link: '/en/guide/' },
    ],
};


module.exports = {
    plugins: [
        searchPlugin(),
        nprogressPlugin(),
    ],

    base: '/go-next-docs/',
    title: 'Next',
    //lang: 'ru-RU',
    description: 'Документация по Next',

    locales: {
        '/ru/': {
            lang: 'ru',
            head: [
                // ['link', { rel: 'icon', href: '/images/logo.png' }], // custom style inject
            ],
        },
        '/en/': {
            lang: 'en',
        },
    },

    theme: defaultTheme({
        search: true,

        locales: {
            '/': ruThemeConfig,
            '/ru/': ruThemeConfig,
            '/en/': enThemeConfig,
        },

        repo: 'alexpts/go-next',

        docsRepo: 'alexpts/go-next-docs/',
        docsBranch: 'master',
        docsDir: 'src',
    }),
}