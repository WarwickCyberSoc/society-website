import defaultImage from './assets/images/default.png';

const CONFIG = {
  name: 'Warwick CyberSoc',

  origin: 'https://society.cyber.warwick.ac.uk',
  basePathname: '/',
  trailingSlash: false,

  title: 'University of Warwick Cyber Security Society',
  description:
    'The University of Warwick Cyber Security Society is a student-run society that aims to promote the study of cyber security and related topics to students at the University of Warwick.',
  defaultImage: defaultImage,

  defaultTheme: 'dark', // Values: "system" | "light" | "dark" | "light:only" | "dark:only"

  language: 'en',
  textDirection: 'ltr',

  dateFormatter: new Intl.DateTimeFormat('en', {
    year: 'numeric',
    month: 'short',
    day: 'numeric',
    timeZone: 'UTC',
  }),

  googleAnalyticsId: false, // or "G-XXXXXXXXXX",
  googleSiteVerificationId: 'orcPxI47GSa-cRvY11tUe6iGg2IO_RPvnA1q95iEM3M',

  blog: {
    disabled: false,
    postsPerPage: 4,

    post: {
      permalink: '/%slug%', // Variables: %slug%, %year%, %month%, %day%, %hour%, %minute%, %second%, %category%
      noindex: false,
      disabled: false,
    },

    list: {
      pathname: 'announcements', // Blog main path, you can change this to "articles" (/articles)
      noindex: false,
      disabled: false,
    },

    tag: {
      pathname: 'tag', // Tag main path /tag/some-tag
      noindex: true,
      disabled: false,
    },
  },
};

export const SITE = { ...CONFIG, blog: undefined };
export const BLOG = CONFIG.blog;
export const DATE_FORMATTER = CONFIG.dateFormatter;
