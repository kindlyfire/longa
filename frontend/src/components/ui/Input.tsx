import React from 'react'
import clsx from 'clsx'

interface InputProps {
	label?: string | JSX.Element
	labelClassName?: string

	type?: React.HTMLInputTypeAttribute
	value?: [string, (v: string) => any]
	placeholder?: string
	disabled?: boolean
	inputClassName?: string
}
export function Input(props: InputProps) {
	return (
		<>
			{props.label && (
				<label className={clsx('block', props.labelClassName)}>
					{props.label}
				</label>
			)}
			<input
				type={props.type || 'text'}
				value={props.value?.[0]}
				placeholder={props.placeholder}
				onInput={e => props.value?.[1](e.currentTarget.value)}
				disabled={props.disabled}
				className={clsx(
					'w-full border border-gray-200 shadow-sm rounded p-2',
					props.inputClassName
				)}
			/>
		</>
	)
}
