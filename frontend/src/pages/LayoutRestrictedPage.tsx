import { Outlet, useNavigate } from 'react-router-dom'
import { useQuery } from 'react-query'
import { Api } from '../state/api'
import { userState } from '../state/user'
import { useEffect } from 'react'

export function LayoutRestrictedPage() {
	const navigate = useNavigate()
	const q = useQuery(
		['userLoggedIn', userState.username, userState.password],
		async () => {
			if (!userState.username || !userState.password) return false
			return Api.checkAuth(userState.username, userState.password).catch(
				e => false
			)
		},
		{
			retry: false,
			refetchOnWindowFocus: false
		}
	)

	useEffect(() => {
		if (q.data === false) navigate('/login')
	}, [q.data])

	if (q.isLoading) return <p>Loading...</p>
	if (q.data === false) return null

	return (
		<div className="container max-w-6xl mx-auto px-2">
			<h1 className="font-bold text-3xl my-4">Longa</h1>

			<Outlet />
		</div>
	)
}
