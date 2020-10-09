import React from 'react'
import { v4 as uuid } from 'uuid'
import Action, { TYPE } from './Action'
import Message from '../Message'

export interface AppState {
  id: string,
  rooms: Array<string>
  messages: {
    [key: string]: Array<Message>
  }
}

export const initialState: AppState = {
  id: uuid(),
  rooms: ['global'],
  messages: {
    global: []
  }
}

const reducer: React.Reducer<AppState, Action> = (state = initialState, action) => {
  switch (action.type) {

    case TYPE.RECEIVE_MESSAGE:
    case TYPE.ADD_MESSAGE:
      const message: Message = action.payload as Message

      return {
        ...state,
        messages: {
          ...state.messages,
          [message.room]: [...state.messages[message.room], message]
        }
      }

    default:
      return state

  }
}

export default reducer
