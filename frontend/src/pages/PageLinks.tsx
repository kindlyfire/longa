import dayjs from 'dayjs'
import { useState } from 'react'
import DataTable from 'react-data-table-component'
import { useMutation, useQuery } from 'react-query'
import { Button } from '../components/ui/Button'
import { Input } from '../components/ui/Input'
import { Api } from '../state/api'
import { Link } from '../state/api-types'
import { queryClient } from '../state/query-client'

const columns = [
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
		name: 'Long URL',
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

export function PageLinks() {
	return (
		<>
			<AddLinksForm />
			<LinksTable />
		</>
	)
}

function AddLinksForm() {
	const [longUrl, setLongUrl] = useState('')
	const [shortUrl, setShortUrl] = useState('')

	const create = useMutation('linkCreate', async () => {
		await Api.linksCreate({ link: shortUrl, target: longUrl })
		queryClient.invalidateQueries('links')
		setLongUrl('')
		setShortUrl('')
	})

	return (
		<>
			{create.isError && (
				<p className="mb-4 bg-red-600 text-white rounded p-2 px-3 shadow font-semibold">
					{'' + create.error}
				</p>
			)}

			<div className="md:grid grid-cols-6 gap-4 mb-4">
				<div className="col-span-3">
					<Input label="Long URL" value={[longUrl, setLongUrl]} />
				</div>
				<div className="col-span-2 mt-4 md:mt-0">
					<Input label="Short URL" value={[shortUrl, setShortUrl]} />
				</div>
				<div className="flex items-end mt-4 md:mt-0">
					<Button
						className="w-full"
						disabled={longUrl === '' || shortUrl === '' || create.isLoading}
						onClick={() => create.mutate()}
					>
						Add
					</Button>
				</div>
			</div>
		</>
	)
}

function LinksTable() {
	const [page, setPage] = useState(0)
	const [selectedRows, setSelectedRows] = useState([] as Link[])
	const [toggleCleared, setToggleCleared] = useState(false)
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

	return (
		<>
			<div className="flex items-center mb-2">
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
			</div>
			<DataTable
				columns={columns}
				data={links.data?.links || []}
				dense
				selectableRows
				selectableRowsVisibleOnly
				onSelectedRowsChange={d => setSelectedRows(d.selectedRows)}
				clearSelectedRows={toggleCleared}
				pagination
				paginationTotalRows={links.data?.count}
				paginationRowsPerPageOptions={[15, 30, 50, 100, 250, 500]}
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
