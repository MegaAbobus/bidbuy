import logo from "../../assets/logo.png";
import style from "./header.module.css";
import { Link } from "react-router-dom";

export const Header = () => {
  return (
    <div className={style.contener}>
      <div className={style.content}>
        <div className={style.link}>
          <Link to={"/"}>
            <img src={logo} alt="" />
          </Link>
          <div className={style.num}>Москва +7 (965) 219-24-67</div>
        </div>
        <nav className={style.navbar}>
          <li>
            <Link className={style.color} to={"/"}>
              Акции
            </Link>
          </li>
          <li>
            <Link className={style.color} to={"/"}>
              Журнал
            </Link>
          </li>
          <li>
            <Link className={style.color} to={"/"}>
              Доставка
            </Link>
          </li>
          <li>
            <Link className={style.color} to={"/"}>
              Обратная связь
            </Link>
          </li>
        </nav>
      </div>
    </div>
  );
};
