import React from 'react'

import { v4 as uuid } from 'uuid'
import Action, { TYPE } from './Action'
import Message from '../Message'

export interface AppState {
  id: string,
  rooms: Array<Room>
  messages: {
    [key: string]: Array<Message>
  }
}

export const initialState: AppState = {
  id: uuid(),
  rooms: [{ id: 'global', name: 'global' }],
  messages: {
    global: []
  }
}

const reducer: React.Reducer<AppState, Action> = (state = initialState, action) => {
  switch (action.type) {

    case TYPE.RECEIVE_MESSAGE:
    case TYPE.ADD_MESSAGE: {

      const message: Message = action.payload as Message

      return {
        ...state,
        messages: {
          ...state.messages,
          [message.room]: [...state.messages[message.room], message]
        }
      }
    }

    case TYPE.SET_ROOMS:
      const rooms: Array<Room> = action.payload as Array<Room>

      let messages: { [key: string]: Array<Message> } = {}

      for (let room of rooms) {
        messages[room.id] = state.messages[room.id] ?? []
      }

      return {
        ...state,
        rooms,
        messages,
      }

    default:
      return state

  }
}

export default reducer
