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
            <Link to={"/auction"} className={style.auction}>
              Аукцион
            </Link>
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
