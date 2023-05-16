import { useState } from "react";
import { Link } from "react-router-dom";
import { Messenger } from "../messenger/messenger";
import style from "./auction.module.css";

export const Auction = () => {
  const [products, setProducts] = useState([
    {
      id: 1,
      image:
        "https://cdn.citilink.ru/sW8mxiEBzIlKrMVEnUfnpR7-cMaWzfWpGnq4IBAfsFQ/resizing_type:fit/gravity:sm/width:1200/height:1200/plain/items/1804836_v01_b.jpg",
      name: "Смартфон vivo T1 6/128Gb, таинственная галактика",
      price: 19990,
      description:
        "VIVO T1 — смартфон с AMOLED-дисплеем, аккумулятором ёмкостью 5000 мАч, быстрой зарядкой FlashCharge и процессором Qualcomm Snapdragon 680. Максимум возможностей для работы и отдыха.",
      date: "2023-05-27",
    },
    {
      id: 2,
      image:
        "https://cdn.citilink.ru/IbsUkzj9ns2IULnGClAOHPbhgKNC-2PmQYrAg9MG3Ds/resizing_type:fit/gravity:sm/width:1200/height:1200/plain/items/1892210_v01_b.jpg",
      name: "Смартфон Huawei nova 10 SE 8/128Gb, серебристый",
      price: 23990,
      description:
        "Смартфон Huawei nova 10 SE — функциональное устройство в тонком корпусе, которое удобно лежит в ладони. Матовое покрытие задней панели не оставляет отпечатков пальцев",
      date: "2023-05-28",
    },
    {
      id: 3,
      image:
        "https://cdn.citilink.ru/E-xkOEpluxpGKlX8mwKc48guNrMjt0D18M40D4Y2kOo/resizing_type:fit/gravity:sm/width:1200/height:1200/plain/items/1782601_v01_b.jpg",
      name: "Смартфон Samsung Galaxy A52s 8/256Gb, SM-A528B, черный",
      price: 28990,
      description:
        "Смартфон Samsung Galaxy A52s выполнен во влаго- и пыленепроницаемом корпусе. Устройство способно выдержать погружение на 1 метр в пресную воду и пробыть там до получаса",
      date: "2023-06-06",
    },
    {
      id: 4,
      image:
        "https://cdn.citilink.ru/2G_Ykw3vOi2xmaowvxPZy8v6rkeQkonAjIS33TveVus/resizing_type:fit/gravity:sm/width:1200/height:1200/plain/items/1841083_v01_b.jpg",
      name: "Смартфон vivo Y35 4/128Gb, черный агат",
      price: 14990,
      description:
        "Смартфон VIVO Y35 выполнен в корпусе, устойчивом к царапинам и интенсивному использованию. Восьмиядерный процессор обеспечивает высокую производительность, плавную и быструю работу приложений.",
      date: "2023-06-01",
    },
    {
      id: 5,
      image:
        "https://cdn.citilink.ru/1pr9zJ1ttRZPVEqwqgDITaeQD4i9fgGPBd_cQ_HSHro/resizing_type:fit/gravity:sm/width:1200/height:1200/plain/items/1823504_v01_b.jpg",
      name: "Смартфон Xiaomi Poco C40 4/64Gb, черный",
      price: 7990,
      description:
        "Смартфон Xiaomi Poco C40 имеет 6,71-дюймовый дисплей Dot Drop с визуальным 2,5D-эффектом и максимальным разрешением 1650×720 пикселей. Экран обладает широкими углами обзора",
      date: "2023-05-21",
    },
    {
      id: 6,
      image:
        "https://cdn.citilink.ru/mP1umTanD3u6K77usjJJL_n75vEmQIvTgS_Ctc4iAWs/resizing_type:fit/gravity:sm/width:1200/height:1200/plain/items/1402188_v01_b.jpg",
      name: "Смартфон Xiaomi Redmi 9A 2/32Gb, синий",
      price: 5990,
      description:
        "Смартфон Xiaomi Redmi 9A – бюджетная модель, подходящая для использования мессенджеров и интернета.",
      date: "2023-06-07",
    },
    {
      id: 7,
      image:
        "https://cdn.citilink.ru/zlkju8NfWkURiSrPseHGeGBMidZvAV2COefYzlEHCwA/resizing_type:fit/gravity:sm/width:1200/height:1200/plain/items/1869931_v01_b.jpg",
      name: "Смартфон REALME 10 4G 8/256Gb, RMX3630, черный",
      price: 21990,
      description:
        "Шустрый, возможно удвоить ОЗУ за счёт ПЗУ. Приятно лежит в руке, относительно небольшой. Неплохой объём памяти.",
      date: "2023-05-20",
    },
    {
      id: 8,
      image:
        "https://cdn.citilink.ru/OF00BK8O5MFbc8Y6JBGpmTleT6WZARvOuMbz8kovRbA/resizing_type:fit/gravity:sm/width:1200/height:1200/plain/items/1804848_v01_b.jpg",
      name: "Смартфон vivo T1 6/128Gb, звездный путь",
      price: 19990,
      description:
        "VIVO T1 — смартфон с AMOLED-дисплеем, аккумулятором ёмкостью 5000 мАч, быстрой зарядкой FlashCharge и процессором Qualcomm Snapdragon 680. Максимум возможностей для работы и отдыха.",
      date: "2023-06-03",
    },
  ]);

  return (
    <div className={style.contant}>
      <div className={style.title}>
        <h2>Лоты</h2>
      </div>
      <Messenger />
      <div className={style.full}>
        <div className={style.cont}>
          <h3>Категории</h3>
          <div>
            <div className={style.titlecategory}>Цена</div>
            <div className={style.inp}>
              <input type="number" placeholder="от 0" />
              <input type="number" placeholder="до 1000000" />
            </div>
          </div>
          <div className={style.category}>
            <div className={style.titlecategory}>Смартфоны и фототехника</div>
            <div className={style.check}>
              <input type="checkbox" />
              <label>Планшеты</label>
            </div>
            <div className={style.check}>
              <input type="checkbox" />
              <label>Смартфоны и гаджеты</label>
            </div>
            <div className={style.check}>
              <input type="checkbox" />
              <label>Фототехника</label>
            </div>
          </div>
          <div className={style.category}>
            <div className={style.titlecategory}>ПК, ноутбуки, переферия</div>
            <div className={style.check}>
              <input type="checkbox" />
              <label>Компьютеры и ПО</label>
            </div>
            <div className={style.check}>
              <input type="checkbox" />
              <label>Ноутбуки и аксессуары</label>
            </div>
            <div className={style.check}>
              <input type="checkbox" />
              <label>Переферия и аксессуары</label>
            </div>
          </div>
          <div className={style.category}>
            <div className={style.titlecategory}>ТВ, консоли и аудио</div>
            <div className={style.check}>
              <input type="checkbox" />
              <label>Аудеотехника</label>
            </div>
            <div className={style.check}>
              <input type="checkbox" />
              <label>Консоли</label>
            </div>
            <div className={style.check}>
              <input type="checkbox" />
              <label>Телевизоры и аксессуары</label>
            </div>
          </div>
          <div className={style.category}>
            <div className={style.titlecategory}>Бытовая техника</div>
            <div className={style.check}>
              <input type="checkbox" />
              <label>Встраимовая техника</label>
            </div>
            <div className={style.check}>
              <input type="checkbox" />
              <label>Техника для дома</label>
            </div>
            <div className={style.check}>
              <input type="checkbox" />
              <label>Техника для кухни</label>
            </div>
          </div>
        </div>
        <div className={style.contener}>
          {products.map((product) => (
            <div className={style.product} key={product.id}>
              <Link className={style.card} to={"#"}>
                <div className={style.info}>
                  <div className={style.photo}>
                    <img src={product.image} alt="" />
                  </div>
                  <div className={style.discrip}>
                    <div className={style.name}>{product.name}</div>
                    <div className={style.detail}>{product.description}</div>
                  </div>
                </div>
                <div className={style.auction}>
                  <div className={style.price}>
                    {product.price.toLocaleString("ru-RU", {
                      currency: "RUB",
                      style: "currency",
                      maximumFractionDigits: 0,
                    })}
                  </div>
                  <div className={style.date}>
                    <span className={style.end}>
                      Дата завершения аукциона: <span>{product.date}</span>
                    </span>
                  </div>
                  <div className={style.btn}>
                    <button>Подробнее</button>
                  </div>
                </div>
              </Link>
            </div>
          ))}
        </div>
      </div>
    </div>
  );
};
