import axios, { AxiosResponse, Method } from 'axios'
import { userState } from './user'

const r = axios.create({
	baseURL: '/_/api/'
})

function handleResponse(d: AxiosResponse) {
	if (d.status !== 200) {
		console.log('API Error:', d.status, d.data)
		const e: any = new Error(
			d.data?.message || 'Non-200 status code API response'
		)
		e.status = d.status
		throw e
	}
	return d.data?.data ?? d.data
}

export const ApiRaw = {
	r,

	async get<T = any>(url: string, query?: object): Promise<T> {
		return ApiRaw.request<T>('GET', url, query)
	},

	async post<T = any>(url: string, body?: object, query?: object): Promise<T> {
		return ApiRaw.request<T>('POST', url, query, body)
	},

	async delete<T = any>(
		url: string,
		body?: object,
		query?: object
	): Promise<T> {
		return ApiRaw.request<T>('DELETE', url, query, body)
	},

	async request<T = any>(
		method: Method,
		url: string,
		query?: any,
		body?: any
	): Promise<T> {
		return r
			.request({
				method,
				url,
				data: body,
				params: query,
				auth: userState.username
					? {
							username: userState.username!,
							password: userState.password!
					  }
					: undefined,
				validateStatus: n => n < 500
			})
			.then(handleResponse)
	}
}
