import { useState } from "react";
import { Link } from "react-router-dom";
import { Messenger } from "../messenger/messenger";
import style from "./auction.module.css";

export const Auction = () => {
  const [products, setProducts] = useState([
    {
      id: 1,
      image:
        "https://cdn.citilink.ru/gk-bP6N-_JHAsclIkhZqaQ7GvvAgZGs_hm3T_g5jbpw/resizing_type:fit/gravity:sm/width:1200/height:1200/plain/items/1892177_v01_b.jpg",
      name: "Смартфон Huawei Mate 50 Pro 8/512Gb",
      description:
        "Смартфон Huawei Mate 50 Pro выполен в корпусе, защищенном от попадания частиц пыли и влаги. Экран устойчив к повреждениям благодаря прочному стеклу Kunlun.",
      price: 79990,
      date: "2023-05-05"
    },
    {
      id: 2,
      image:
        "https://cdn.citilink.ru/k7Tw00w5zTyKad7vouI1goF6u2NMyo5Vq1Xf3nqsNps/resizing_type:fit/gravity:sm/width:1200/height:1200/plain/items/1800332_v01_b.jpg",
      name: "Смартфон TECNO Camon 19 6/128Gb, черный",
      price: 9990,
    },
    {
      id: 3,
      image:
        "https://cdn.citilink.ru/9gMoiNKDS6Nf-KO9K_8d60pYv3XalmTXygklmnq7AuY/resizing_type:fit/gravity:sm/width:1200/height:1200/plain/items/1913377_v01_b.jpg",
      name: "Смартфон vivo V27e 8/256Gb, черный оникс",
      price: 26990,
    },
    {
      id: 4,
      image:
        "https://cdn.citilink.ru/ux3xZ82_UueNv5st8JeiZvgIO7v1umgdkz59xkCkoRw/resizing_type:fit/gravity:sm/width:1200/height:1200/plain/items/1594036_v01_b.jpg",
      name: "Смартфон REALME C25s 4/128Gb, голубой",
      price: 18550,
    },
    {
      id: 5,
      image:
        "https://cdn.citilink.ru/juoffa0c9DUbCn7FpHJXHn-UwSXWTQxTVaCH4z5Kf4w/resizing_type:fit/gravity:sm/width:1200/height:1200/plain/items/1783104_v01_b.jpg",
      name: "Смартфон REALME C35 4/128Gb, RMX3511, черный",
      price: 12990,
    },
    {
      id: 6,
      image:
        "https://cdn.citilink.ru/_-7uaCcsJcXVfh7JYZPqUpfWOLnhGajEF70NQ0Gkx_4/resizing_type:fit/gravity:sm/width:1200/height:1200/plain/items/1804652_v01_b.jpg",
      name: "Смартфон OPPO A96 6/128Gb, черный",
      price: 16990,
    },
    {
      id: 7,
      image:
        "https://cdn.citilink.ru/G2SbWZZbasNvInlrdEsyhEjGWai65zvoksH8Bb12cRc/resizing_type:fit/gravity:sm/width:1200/height:1200/plain/items/1628173_v01_b.jpg",
      name: "Смартфон REALME C25s 4/64Gb, серый",
      price: 8990,
    },
    {
      id: 8,
      image:
        "https://cdn.citilink.ru/KqL3YroNV7LK9ROi_eWfd1EzR0l93yN7dN1F5CEy9u4/resizing_type:fit/gravity:sm/width:1200/height:1200/plain/items/1517181_v01_b.jpg",
      name: "Смартфон REALME 8 6/128Gb, черный",
      price: 19990,
    },
  ]);

  return (
    <div className={style.contant}>
      <div className={style.title}>
        <h2>Лоты</h2>
      </div>
      <Messenger />
      {products.map((product) => (
        <div className={style.product} key={product.id}>
          <Link className={style.card} to={"/"}>
            <div className={style.info}>
              <div className={style.photo}>
                <img src={product.image} alt="" />
              </div>
              <div className={style.discrip}>
                <div className={style.name}>{product.name}</div>
                <div>{product.description}</div>
              </div>
            </div>
            <div className={style.price}>
              {product.price.toLocaleString("ru-RU", {
                currency: "RUB",
                style: "currency",
                maximumFractionDigits: 0,
              })}
            <div>{product.date}</div>
            </div>
          </Link>
        </div>
      ))}
    </div>
  );
};
