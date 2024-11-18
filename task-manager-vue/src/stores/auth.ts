import AuthService from "@/services/AuthService";
import { AuthSignIn, AuthSignUp, UserSignIn } from "@/types/Auth";
import { AxiosResponse } from "axios";
import { defineStore } from "pinia";
import { Response } from "@/types/Response";
import { toast } from "vue3-toastify";
import { AUTH } from "@/Constants";
import { getCookie, removeCookie, setCookie } from "@/utils/jsCookie";

export const useAuthStore = defineStore("auth", () => {
  interface AuthStateType {
    loading: boolean;
    userSignIn: UserSignIn;
  }

  const initState: AuthStateType = {
    loading: false,
    userSignIn: {
      user_id: "",
      username: "",
      token: "",
    },
  };

  const router = useRouter();
  const auth = reactive<AuthStateType>(initState);

  const alertErrorMessageSignIn = (statusCode: number, message?: string) => {
    switch (statusCode) {
      case 500:
        toast.error("เกิดข้อผิดพลาดในการเข้าสู่ระบบ");
        break;

      case 422:
        toast.error("กรุณากรอกชื่อผู้ใช้งานให้ถูกต้อง");
        break;

      case 400:
        toast.error("กรุณากรอกข้อมูลให้ถูกต้อง");
        break;

      case 403:
        if (message === "Account is locked.") {
          toast.error(
            `บัญชีผู้ใช้งานถูกระงับการใช้งาน \nเนื่องจากกรอกรหัสผ่านไม่ถูกต้องเกิน 5 ครั้ง`,
            {
              position: toast.POSITION.TOP_CENTER,
            }
          );
          setTimeout(() => {
            toast.warning("กรุณาติดต่อผู้ดูแลระบบ", {
              position: toast.POSITION.TOP_CENTER,
            });
          }, 1300);
          break;
        }

        toast.error(message);
        break;

      case 401:
        let errorMessage = "กรุณากรอกรหัสผ่านให้ถูกต้อง";
        if (message === "Invalid OTP code.") {
          errorMessage = "กรุณากรอกรหัส 2fa ให้ถูกต้อง";
        }
        toast.error(errorMessage);
        break;

      default:
        toast.error("กรุณากรอกข้อมูลให้ถูกต้อง");
        break;
    }
  };

  const alertErrorMessageSignUp = (statusCode: number, message?: string) => {
    switch (statusCode) {
      case 500:
        toast.error("เกิดข้อผิดพลาดในการลงทะเบียน");
        break;

      case 422:
        toast.error("กรุณากรอกข้อมูลให้ถูกต้อง");
        break;

      case 400:
        let mess = "กรุณากรอกข้อมูลให้ถูกต้อง";
        if (message === "มีชื่อบัญชีผู้ใช้งานนี้เเล้ว") {
          mess = message;
        }
        toast.error(mess);
        break;

      default:
        toast.error("กรุณากรอกข้อมูลให้ถูกต้อง");
        break;
    }
  };

  const isSignInUserAuth = (): boolean => {
    const userSignIn = getCookie(AUTH.userSignIn);
    if (!userSignIn) return false;
    auth.userSignIn = JSON.parse(userSignIn);
    return true;
  };

  const handleSignIn = async (payload: AuthSignIn) => {
    auth.loading = true;
    try {
      const { data }: AxiosResponse = await AuthService.signIn(payload);
      const { result }: Response = data;

      const userSignIn: UserSignIn = {
        ...result,
      };

      toast.success("เข้าสู่ระบบสําเร็จ");
      auth.userSignIn = userSignIn;
      setCookie(AUTH.userSignIn, userSignIn);
      router.replace("/");
      return userSignIn;
    } catch (error: any) {
      console.log("error", error);
      const { message }: any = error.response.data;
      console.error(error);
      alertErrorMessageSignIn(error.status, message);
    } finally {
      auth.loading = false;
    }
  };

  const handleSignUp = async (payload: AuthSignUp) => {
    auth.loading = true;
    try {
      const { data }: AxiosResponse = await AuthService.signUp(payload);
      const { result }: Response = data;

      toast.success("ลงทะเบียนสำเร็จ");
      setTimeout(() => {
        router.replace(`/auth/regis2Fa?username=${payload.username}`);
      }, 1500);
      return result;
    } catch (error: any) {
      console.log("error", error);
      const { message }: any = error.response.data;
      console.error(error);
      alertErrorMessageSignUp(error.status, message);
    } finally {
      auth.loading = false;
    }
  };

  const handleSignOut = async () => {
    auth.loading = true;
    try {
      await removeCookie(AUTH.userSignIn);
      await router.replace("/auth/signIn");
    } catch (error: any) {
      toast.error("เกิดข้อผิดพลาดในการออกจากระบบ");
      console.error(error);
    } finally {
      auth.loading = false;
    }
  };

  const findUserSignInById = async (username: string) => {
    try {
      const { data }: AxiosResponse = await AuthService.findUserSignInById(
        username
      );
      const { result }: Response = data;
      return result;
    } catch (error: any) {
      console.log("error", error);
    }
  };

  return {
    // export state
    auth,
    // export actions
    handleSignIn,
    handleSignUp,
    handleSignOut,
    isSignInUserAuth,
    findUserSignInById,
  };
});
