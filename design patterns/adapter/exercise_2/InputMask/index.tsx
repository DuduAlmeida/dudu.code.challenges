import { forwardRef } from "react";
import { InputMaskProps } from "./types";

import inputMaskAdapter from "./adapters/ReactImask";

// eslint-disable-next-line react/display-name, no-undef
const InputWithoutMask = forwardRef<HTMLInputElement, any>((props, ref) => (
  <input {...props} />
));

const InputMask = ({ mask, ...props }: InputMaskProps) => {
  if (!inputMaskAdapter.isValid(mask)) return <InputWithoutMask {...props} />;

  return <inputMaskAdapter.Component mask={mask} {...props} />;
};

export default InputMask;
