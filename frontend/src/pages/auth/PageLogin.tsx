import { runInAction } from 'mobx'
import { observer } from 'mobx-react-lite'
import { useState } from 'react'
import { useMutation } from 'react-query'
import { useNavigate } from 'react-router-dom'
import { Button } from '../../components/ui/Button'
import { Input } from '../../components/ui/Input'
import { Api } from '../../state/api'
import { userState } from '../../state/user'

export const PageLogin = observer(function PageLogin() {
	const navigate = useNavigate()
	const [username, setUsername] = useState('')
	const [password, setPassword] = useState('')

	const login = useMutation('login', async () => {
		await Api.checkAuth(username, password)
		runInAction(() => {
			userState.username = username
			userState.password = password
		})
		navigate('/')
	})

	return (
		<div className="mx-auto container max-w-md sm:border border-gray-100 sm:shadow-md rounded-lg p-4 mt-8">
			<h1 className="font-bold text-3xl text-center mb-4">Longa Admin</h1>

			{login.isError && (
				<p className="mb-4 bg-red-600 text-white rounded p-2 px-3 shadow font-semibold">
					{'' + login.error}
				</p>
			)}

			<Input
				label="Username"
				type="text"
				value={[username, setUsername]}
				inputClassName="mb-4"
			/>
			<Input
				label="Password"
				type="password"
				value={[password, setPassword]}
				inputClassName="mb-4"
			/>

			<div className="flex justify-center">
				<Button disabled={login.isLoading} onClick={() => login.mutate()}>
					Submit
				</Button>
			</div>
		</div>
	)
})
