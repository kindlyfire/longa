import dayjs from 'dayjs'
import { useState } from 'react'
import DataTable, { ExpanderComponentProps } from 'react-data-table-component'
import { useMutation, useQuery } from 'react-query'
import { Button } from '../components/ui/Button'
import { Input } from '../components/ui/Input'
import { Api } from '../state/api'
import { Link } from '../state/api-types'
import { queryClient } from '../state/query-client'

export function PageLinks() {
	return (
		<>
			<h2 className="font-bold text-xl mb-4">Create</h2>
			<AddLinksForm />

			<h2 className="font-bold text-xl mb-4">Links</h2>
			<LinksTable />
		</>
	)
}

function AddLinksForm() {
	const [targetUrl, setTargetUrl] = useState('')
	const [shortUrl, setShortUrl] = useState('')

	const create = useMutation('linkCreate', async () => {
		await Api.linksCreate({ link: shortUrl, target: targetUrl })
		queryClient.invalidateQueries('links')
		setTargetUrl('')
		setShortUrl('')
	})

	return (
		<>
			{create.isError && (
				<p className="mb-4 bg-red-600 text-white rounded p-2 px-3 shadow font-semibold">
					{'' + create.error}
				</p>
			)}

			<form
				onSubmit={e => {
					e.preventDefault()
					create.mutate()
				}}
			>
				<div className="md:grid grid-cols-6 gap-4 mb-4">
					<div className="col-span-3">
						<Input label="Target URL" value={[targetUrl, setTargetUrl]} />
					</div>
					<div className="col-span-2 mt-4 md:mt-0">
						<Input label="Short URL" value={[shortUrl, setShortUrl]} />
					</div>
					<div className="flex items-end mt-4 md:mt-0">
						<Button
							className="w-full"
							disabled={targetUrl === '' || shortUrl === '' || create.isLoading}
						>
							Add
						</Button>
					</div>
				</div>
			</form>
		</>
	)
}

const linksTableColumns = [
	{
		name: 'Short URL',
		cell: (row: Link) => (
			<a
				className="underline whitespace-nowrap overflow-hidden text-ellipsis"
				href={'/' + row.link}
				target="_blank"
			>
				{row.link}
			</a>
		)
	},
	{
		name: 'Target URL',
		selector: (row: Link) => row.target,
		cell: (row: Link) => (
			<a
				className="underline whitespace-nowrap overflow-hidden text-ellipsis"
				href={row.target}
				target="_blank"
			>
				{row.target}
			</a>
		)
	},
	{
		name: 'Date',
		sortable: true,
		selector: (row: Link) => dayjs(row.createdAt).format('DD/MM/YYYY HH:mm:ss')
	}
]

function LinksTable() {
	// page is unused for now
	const [page, setPage] = useState(0)
	const [selectedRows, setSelectedRows] = useState([] as Link[])
	const [toggleCleared, setToggleCleared] = useState(false)
	const [searchString, setSearchString] = useState('')
	const links = useQuery(['links', page], () =>
		Api.linksList({ offset: page * 15 })
	)

	const del = useMutation('linksDelete', async () => {
		// This would need a direct way to delete multiple links on the server
		// to be efficient, which we don't have, and meh
		for (let row of selectedRows) {
			await Api.linksDelete(row.id)
		}
		queryClient.invalidateQueries('links')
		setToggleCleared(!toggleCleared)
		setSelectedRows([])
	})

	const displayLinks =
		(searchString
			? links.data?.links.filter(
					l => l.link.includes(searchString) || l.target.includes(searchString)
			  )
			: links.data?.links) ?? []

	return (
		<>
			<div className="flex items-center mb-4">
				<Button
					color="red"
					size="sm"
					className="mr-3"
					disabled={selectedRows.length === 0}
					onClick={() => del.mutate()}
				>
					Delete
				</Button>
				<div>
					{selectedRows.length} row{selectedRows.length !== 1 && 's'} selected
				</div>
				<div className="ml-auto">
					<Input
						placeholder="Search"
						inputClassName="py-1"
						value={[searchString, setSearchString]}
					/>
				</div>
			</div>
			<DataTable
				columns={linksTableColumns}
				data={displayLinks}
				dense
				selectableRows
				selectableRowsVisibleOnly
				onSelectedRowsChange={d => setSelectedRows(d.selectedRows)}
				clearSelectedRows={toggleCleared}
				pagination
				paginationTotalRows={links.data?.count}
				paginationRowsPerPageOptions={[15, 30, 50, 100, 250, 500]}
				expandableRows
				expandableRowsComponent={LinksTableEdit}
				customStyles={{
					rows: {
						denseStyle: {
							fontSize: 'inherit'
						}
					},
					head: {
						style: {
							fontSize: '0.9rem'
						}
					}
				}}
			></DataTable>
		</>
	)
}

function LinksTableEdit(props: ExpanderComponentProps<Link>) {
	const [targetUrl, setTargetUrl] = useState(props.data.target)
	const [shortUrl, setShortUrl] = useState(props.data.link)

	const update = useMutation('linkUpdate', async () => {
		await Api.linksUpdate({ ...props.data, link: shortUrl, target: targetUrl })
		queryClient.invalidateQueries('links')
	})

	return (
		<div className="m-2">
			{update.isError && (
				<p className="mb-4 bg-red-600 text-white rounded p-2 px-3 shadow font-semibold">
					{'' + update.error}
				</p>
			)}

			<form
				onSubmit={e => {
					e.preventDefault()
					update.mutate()
				}}
			>
				<div className="md:grid grid-cols-6 gap-4 mb-4">
					<div className="col-span-3">
						<Input label="Target URL" value={[targetUrl, setTargetUrl]} />
					</div>
					<div className="col-span-2 mt-4 md:mt-0">
						<Input label="Short URL" value={[shortUrl, setShortUrl]} />
					</div>
					<div className="flex items-end mt-4 md:mt-0">
						<Button
							className="w-full"
							disabled={
								targetUrl === '' ||
								shortUrl === '' ||
								update.isLoading ||
								(targetUrl === props.data.target &&
									shortUrl === props.data.link)
							}
						>
							Save
						</Button>
					</div>
				</div>
			</form>
		</div>
	)
}
