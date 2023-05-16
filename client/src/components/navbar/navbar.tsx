import {
  faBars,
  faCartPlus,
  faMagnifyingGlass,
  faUser,
} from "@fortawesome/free-solid-svg-icons";
import { FontAwesomeIcon } from "@fortawesome/react-fontawesome";
import style from "./navbar.module.css";
import { Link } from "react-router-dom";

export const NavBar = () => {
  return (
    <div className={style.contener}>
      <div className={style.contant}>
        <div className={style.btnserch}>
          <div className={style.btn}>
            <button>
              <FontAwesomeIcon icon={faBars} className={style.menu} />
              Каталог товаров
            </button>
          </div>
          <div className={style.serch}>
            <form className={style.sub} action="" method="get">
              <input name="s" placeholder="Поиск по товарам" type="search" />
              <button type="submit">
                <FontAwesomeIcon icon={faMagnifyingGlass} />
              </button>
            </form>
          </div>
        </div>
        <div className={style.nav}>
          <Link className={style.sign} to={"/signin"}>
            <FontAwesomeIcon icon={faUser} />
            <span>Войти</span>
          </Link>
          <Link className={style.sign} to={"/"}>
            <FontAwesomeIcon icon={faCartPlus} />
            <span>Корзина</span>
          </Link>
        </div>
      </div>
    </div>
  );
};

{
  /* <ul>
  <li>
    <a href="">Бытовая техника</a>
  </li>
  <li>
    <a href="">Смартфоны</a>
  </li>
  <li>
    <a href="">Смартфоны</a>
  </li>
  <li>
    <a href="">Фототехника</a>
  </li>
  <li>
    <a href="">ТВ</a>
  </li>
  <li>
    <a href="">Консоли</a>
  </li>
  <li>
    <a href="">Пк</a>
  </li>
  <li>
    <a href="">Ноутбуки</a>
  </li>
  <li>
    <a href="">Периферия</a>
  </li>
  <li>
    <a href="">Сетевое оборудование</a>
  </li>
</ul> */
}
