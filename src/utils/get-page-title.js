import defaultSettings from '@/settings'

const title = defaultSettings.title || 'Go Element Admin'

export default function getPageTitle(pageTitle) {
  if (pageTitle) {
    return `${pageTitle} - ${title}`
  }
  return `${title}`
}
