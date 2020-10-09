import React, { useCallback, useState } from 'react'
import './Input.css'

interface props {
 onSubmit: (value: string) => void
}

const Input: React.FunctionComponent<props> = ({ onSubmit }) => {
  const [message, setMessage] = useState('')

  const onChange = useCallback((event: React.SyntheticEvent<HTMLInputElement>) => {
    setMessage(event.currentTarget.value)
  }, [])

  function onClick () {
    onSubmit(message)
    setMessage('')
  }

  function onKeyDown (event: React.KeyboardEvent<HTMLInputElement>) {
    if (event.key === 'Enter') {
      onSubmit(message)
      setMessage('')
    }
  }

  return (
    <div className="input">
      <input type="text" value={message} onChange={onChange} onKeyDown={onKeyDown} />

      <button onClick={onClick}>
        Submit
      </button>

    </div>
  )

}

export default Input