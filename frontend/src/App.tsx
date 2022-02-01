import { QueryClientProvider } from 'react-query'
import { Routes, Route, HashRouter } from 'react-router-dom'
import './App.css'
import { PageLogin } from './pages/auth/PageLogin'
import { PageLinks } from './pages/PageLinks'
import { LayoutRestrictedPage } from './pages/LayoutRestrictedPage'
import { queryClient } from './state/query-client'
import { userState } from './state/user'

export function App() {
	userState.load()

	return (
		<QueryClientProvider client={queryClient}>
			<HashRouter>
				<Routes>
					<Route path="login" element={<PageLogin />} />

					<Route element={<LayoutRestrictedPage />}>
						<Route index element={<PageLinks />} />
					</Route>
				</Routes>
			</HashRouter>
		</QueryClientProvider>
	)
}
