import { FontAwesomeIcon } from "@fortawesome/react-fontawesome";
import {
  faClose,
  faMessage,
  faUserSecret,
} from "@fortawesome/free-solid-svg-icons";
import { useState } from "react";
import { Dialog } from "@headlessui/react";
import style from "./messenger.module.css";

export const Messenger = () => {
  const [isOpen, setIsOpen] = useState<boolean>(false);

  return (
    <div onClick={() => setIsOpen(true)} className={style.content}>
      <div className={style.mes}>
        <FontAwesomeIcon className={style.icon} icon={faMessage} />
      </div>
      <Dialog open={isOpen} onClose={() => setIsOpen(false)}>
        <div className={style.bg}>
          <Dialog.Panel className={style.popup}>
            <div className={style.btn}>
              <FontAwesomeIcon
                onClick={() => {
                  setIsOpen(false);
                }}
                className={style.btnicon}
                icon={faClose}
              />
            </div>
            <div className={style.head}>
              <div className={style.title}>Служба поддержки</div>
              <span>online</span>
            </div>
            <div className={style.main}>
              <div className={style.hi}>
                <FontAwesomeIcon
                  className={style.assistant}
                  icon={faUserSecret}
                />
                <span>Здравствуйте!👋 У вас появился вопрос?</span>
              </div>
              <div className={style.inp}>
                <input type="text" placeholder="Введите сообщение..." />
              </div>
            </div>
          </Dialog.Panel>
        </div>
      </Dialog>
    </div>
  );
};
