import { Link } from "react-router-dom";
import style from "./signup.module.css";

export const SigUp = () => {
  return (
    <div className={style.contant}>
      <form className={style.contener}>
        <h1 className={style.title}>Регистрация</h1>
        <div className={style.fields}>
          <div className={style.field}>
            <input type="text" placeholder="Ваше имя" />
          </div>
          <div className={style.field}>
            <input type="text" placeholder="Номер телефона" />
          </div>
          <div className={style.field}>
            <input type="password" placeholder="Пароль" />
          </div>
        </div>
        <div className={style.btns}>
          <button className={style.btn_up}>Зарегистрироваться</button>
          <div>Есть аккаунт?</div>
          <Link to={"/signin"} className={style.btn}>
            <button>Войти</button>
          </Link>
        </div>
      </form>
    </div>
  );
};
