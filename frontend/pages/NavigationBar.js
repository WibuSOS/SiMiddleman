import { useState } from "react";
import NavItem from "./NavigationItem";
import Link from "next/link";
import Logo from './assets/logo2.png';

const MENU_LIST = [
  { text: "Home", href: "/" },
  { text: "Customer Service", href: "/customerservice"},
];

const Navbar = () => {
  const [navActive, setNavActive] = useState(null);
  const [activeIdx, setActiveIdx] = useState(-1);

  return (
    <header>
      <nav className='nav'>
        <Link href={"/"}>
          <a className="logo-brand">
            <img src={Logo.src} className='logo-navbar me-3'/>
            SiMiddleman+
          </a>
        </Link>
        <div onClick={() => setNavActive(!navActive)} className='nav__menu-bar'>
          <div></div>
          <div></div>
          <div></div>
        </div>
        <div className={`${navActive ? "active" : ""} nav__menu-list`}>
          {MENU_LIST.map((menu, idx) => {
            return (
              <div className="link-div" key={menu.text} onClick={() => {
                setActiveIdx(idx);
                setNavActive(false);
              }}>
                <NavItem active={activeIdx === idx} {...menu} />
              </div>
            )
          })}
        </div>
      </nav>
    </header>
  );
};

export default Navbar;