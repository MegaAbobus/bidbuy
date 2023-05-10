import { Link } from "react-router-dom";
import style from "./signin.module.css";

export const SigIn = () => {
  return (
    <div className={style.contant}>
      <form className={style.contener}>
        <h1 className={style.title}>Вход</h1>
        <div className={style.fields}>
          <div className={style.field}>
            <input type="text" placeholder="Номер телефона" />
          </div>
          <div className={style.field}>
            <input type="password" placeholder="Пароль" />
          </div>
        </div>
        <div className={style.btns}>
          <button className={style.btn}>Войти</button>
          <div>Не зарегестрированы?</div>
          <Link to={"/signup"} className={style.btn_up}>
            <button>Зарегистрироваться</button>
          </Link>
        </div>
      </form>
    </div>
  );
};
