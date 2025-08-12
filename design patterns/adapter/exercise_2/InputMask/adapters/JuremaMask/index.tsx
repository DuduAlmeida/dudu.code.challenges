import React from "react";

import {
  InputMaskAdapterPayload,
  InputMaskAdapterResponse,
} from "../definition";

const JuremaMask: React.FC<InputMaskAdapterPayload> = (props) => {
  //Lógica do JuremaMask
  return <input {...props} />;
};

const adapter: InputMaskAdapterResponse = {
  Component: JuremaMask,
  isValid: (mask: string) => !mask || mask.startsWith("jurema"),
};
export default adapter;
