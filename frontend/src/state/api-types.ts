interface BaseModel {
	id: number
}

export interface Link extends BaseModel {
	link: string
	target: string
	createdAt: string
}
