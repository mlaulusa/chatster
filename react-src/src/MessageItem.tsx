import React from 'react'
import './MessageItem.css'
import Message from './Message'

interface props {
  value: Message
}

const MessageItem: React.FunctionComponent<props> = ({ value }) => (
  <div className="message">

    <span>{value.user}</span>
    <span>{value.payload}</span>
    <span>{value.time.toFormat('tt')}</span>

  </div>
)

export default MessageItem