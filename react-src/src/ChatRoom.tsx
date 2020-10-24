import React, { useCallback, useContext } from 'react'
import MessageList from './MessageList'
import Input from './Input'
import Creator from './state/Creator'
import { AppContext } from './state/Provider'
import useRoomWebSocket from './hooks/useRoomWebSocket'
import Message from './Message'

interface props {
  name: string
}

const ChatRoom: React.FunctionComponent<props> = ({ name }) => {

  const { state: { id }, dispatch } = useContext(AppContext)

  const send = useRoomWebSocket(name)

  const onSubmit = useCallback((value: string) => {
    const message: Message = new Message(id, value, name)
    send(message)
    dispatch(Creator.receiveMessage(message))
  }, [send])

  return (
    <div>
      <h2>{name}</h2>

      <MessageList room={name} />

      <Input onSubmit={onSubmit} />

    </div>
  )
}

export default ChatRoom
