import React from 'react';


export interface InputMaskAdapterPayload {
  onChange: (event: any) => void
  mask: string
}

export interface InputMaskAdapterResponse {
  isValid: (mask: string) => boolean
  Component: React.FC<InputMaskAdapterPayload>
}