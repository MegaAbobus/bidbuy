import logo from "../../assets/logo.png";
import style from "./footer.module.css";

export const Footer = () => {
  return (
    <footer className={style.colorfooter}>
      <div className={style.logo}>
        <img src={logo} alt="" />
        <p>
          Don't worry if something doesn't work. If everything worked, you'd be
          fired.
        </p>
      </div>
      © Web development 2023
    </footer>
  );
};
