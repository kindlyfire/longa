import { ApiRaw } from './api-raw'
import { Link } from './api-types'

export const Api = {
	checkAuth: (username: string, password: string) =>
		ApiRaw.r
			.get('links', {
				auth: {
					username,
					password
				}
			})
			.then(() => true),

	linksList: (query: { offset: number }) =>
		ApiRaw.get<{
			links: Link[]
			count: number
		}>('links', query),
	linksCreate: (link: Partial<Link>) => ApiRaw.post<Link>('links', link),
	linksUpdate: (link: Link) => ApiRaw.post<Link>('links/' + link.id, link),
	linksDelete: (linkId: number) => ApiRaw.delete<true>('links/' + linkId)
}
