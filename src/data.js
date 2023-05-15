import { getPermalink, getBlogPermalink, getHomePermalink } from './utils/permalinks';

export const headerData = {
  links: [
    {
      text: 'Home',
      href: getHomePermalink(),
    },
    {
      text: 'Announcements',
      href: getBlogPermalink(),
    },
  ],
  actions: [
    { type: 'button', text: 'Join The Society', href: 'https://www.warwicksu.com/societies-sports/societies/61481/' }
  ],
};
  
export const footerData = {
  secondaryLinks: [
    { text: 'Report a Security Issue', href: 'mailto:cybersoc@warwick.ac.uk' },
  ],
  socialLinks: [
    { ariaLabel: 'Instagram', icon: 'tabler:brand-instagram', href: 'https://www.instagram.com/warwickcybersoc/' },
    { ariaLabel: 'Github', icon: 'tabler:brand-github', href: 'https://github.com/wmgcyber' },
    { ariaLabel: 'Discord', icon: 'tabler:brand-discord', href: ' https://discord.com/invite/WkKTxjEzUt'},
    { ariaLable: 'Mail', icon: 'tabler:mail', href: 'mailto:cybersoc@warwick.ac.uk' },
    { type: 'button', text: 'Join The Society', href: 'https://www.warwicksu.com/societies-sports/societies/61481/' },
  ],
  footNote: `
    Theme by <a class="text-blue-600 hover:underline dark:text-gray-200" href="https://onwidget.com/"> onWidget</a> ·  © 2023 Warwick Cyber Security Society.
  `,
};
