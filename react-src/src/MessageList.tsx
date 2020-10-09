import React, { useContext } from 'react'
import map from 'lodash/map'
import MessageItem from './MessageItem'
import { AppContext } from './state/Provider'
import "./MessageList.css"

interface props {
  room: string
}

const MessageList: React.FunctionComponent<props> = ({ room }) => {

  const { state: { messages: { [room]: messages } } } = useContext(AppContext)

  return (
    <div className="message-list">
      {map(messages, message => <MessageItem value={message} />)}
    </div>
  )

}

export default MessageList
// @ts-ignore