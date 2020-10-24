import { useCallback, useContext, useEffect, useState } from 'react'
import Message from '../Message'
import { AppContext } from '../state/Provider'
import Creator from '../state/Creator'

function serializeMessage (message: Message): string {

  return JSON.stringify(message.toJSON())
}

function deserializeMessage (message: string): Message {

  return Message.fromJSON(JSON.parse(message))
}

export default function useRoomWebSocket (room: string = 'global') {

  const { dispatch } = useContext(AppContext)
  const [webSocket, setWebSocket] = useState<WebSocket>()

  useEffect(() => {

    const socket = new WebSocket(`ws://localhost:3000/room/${room}`)

    function onOpen () {
      setWebSocket(socket)
    }

    function onMessage (event: MessageEvent) {
      const message = deserializeMessage(event.data)

      dispatch(Creator.receiveMessage(message))

    }

    socket.addEventListener('open', onOpen)
    socket.addEventListener('message', onMessage)

    return () => {

      webSocket?.removeEventListener('open', onOpen)
      webSocket?.removeEventListener('message', onMessage)
      socket?.close()
    }
  }, [])

  return useCallback((message: Message) => {

    webSocket?.send(serializeMessage(message))

  }, [webSocket])
}
