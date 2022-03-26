import { FC, useState } from "react";

type Props = {
  isOpen: boolean;
  onClose: () => void;
};

export const Dialog: FC<Props> = ({ isOpen, onClose, children }) => {
  //   const [show, setShow] = useState<boolean>(false);
  return (
    <div className={`fixed inset-0 z-30 ${isOpen ? "visible" : "invisible"}`}>
      <div
        className="fixed inset-0"
        onClick={() => {
          onClose();
        }}
      >
        <div className="absolute inset-0 bg-gray-500 opacity-75"></div>
      </div>
      <div className="fixed z-10 bg-white">{children}</div>
    </div>
  );
};
