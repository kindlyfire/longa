import clsx from 'clsx'
import React from 'react'

const colorClasses = {
	lime: 'bg-lime-600 border-lime-600',
	red: 'bg-red-600 border-red-600'
}

const sizeClasses = {
	sm: 'p-1 px-2',
	md: 'p-2 px-3'
}

interface ButtonProps {
	children?: React.ReactNode
	onClick?: () => any
	disabled?: boolean
	className?: string
	color?: keyof typeof colorClasses
	size?: keyof typeof sizeClasses
}
export function Button(props: ButtonProps) {
	return (
		<button
			className={clsx(
				'text-white rounded shadow font-semibold border',
				colorClasses[props.color || 'lime'],
				sizeClasses[props.size || 'md'],
				props.disabled ? 'bg-opacity-70' : 'bg-opacity-90',
				props.className
			)}
			onClick={props.onClick}
			disabled={props.disabled}
		>
			{props.children}
		</button>
	)
}
