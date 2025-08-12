import React from "react";

import {
  InputMaskAdapterPayload,
  InputMaskAdapterResponse,
} from "../definition";

const ReactImask: React.FC<InputMaskAdapterPayload> = (props) => {
  //Lógica do ReactImask
  return <input {...props} />;
};

const adapter: InputMaskAdapterResponse = {
  Component: ReactImask,
  isValid: (mask: string) => !mask || typeof mask === "function",
};
export default adapter;
