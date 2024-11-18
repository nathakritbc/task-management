import Cookies from "js-cookie";

// กำหนดประเภทของชื่อคุกกี้
type CookieName = string;

// กำหนดประเภทของค่าในคุกกี้
type CookieValue = string | object | any;

const getCookie = (name: CookieName): string | undefined => {
  return Cookies.get(name);
};

const getAllCookies = (): { [key: string]: string } => {
  return Cookies.get();
};

const setCookie = (
  name: CookieName,
  value: CookieValue,
  expires: number = 1
): void => {
  const payload = JSON.stringify(value);
  Cookies.set(name, payload, { expires });
};

const setStringCookie = (
  name: CookieName,
  value: string,
  expires: number = 1
): void => {
  Cookies.set(name, value, { expires });
};

const getStringCookie = (name: CookieName): string | undefined => {
  return Cookies.get(name);
};

const removeCookie = (name: CookieName): void => {
  Cookies.remove(name);
};

export {
  getCookie,
  setCookie,
  removeCookie,
  getAllCookies,
  setStringCookie,
  getStringCookie,
};
