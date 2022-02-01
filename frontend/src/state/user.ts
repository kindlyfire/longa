import { observable, reaction } from 'mobx'

const LOCALSTORAGE_KEY = 'longa_tokens'

export const userState = observable({
	username: null as string | null,
	password: null as string | null,

	load() {
		if (userState.username) return

		const payload = localStorage.getItem(LOCALSTORAGE_KEY)
		if (!payload) return

		try {
			const data = JSON.parse(payload)
			userState.username = data.username
			userState.password = data.password
		} catch (e) {
			return
		}
	},
	save() {
		localStorage.setItem(
			LOCALSTORAGE_KEY,
			JSON.stringify({
				username: userState.username,
				password: userState.password
			})
		)
	}
})

reaction(
	() => [userState.username, userState.password],
	() => {
		userState.save()
	}
)
