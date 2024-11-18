<script lang="ts" setup>
import { AxiosResponse } from "axios";
import AuthService from "@/services/AuthService";
import { Response } from "@/types/Response";
import { toast } from "vue3-toastify";
import QrcodeVue from "qrcode.vue";
const route = useRoute();
const router = useRouter();
const qrCodeUrl = ref("");
const user = ref("");
const loading = ref(false);

const handleQrCode2Fa = async (payload: string) => {
  loading.value = true;
  try {
    router.push(`/auth/regis2Fa?username=${payload}`);
    const { data }: AxiosResponse<Response> = await AuthService.enable2FA(
      payload
    );
    if (data.result.qr_url) {
      qrCodeUrl.value = data.result.qr_url;
      return true;
    }
  } catch (error) {
    user.value = "";
    qrCodeUrl.value = "";
    toast.error("ชื่อผู้ใช้งานไม่ถูกต้อง");
    console.error(error);
    setTimeout(() => {
      router.push(`/auth/regis2Fa`);
    }, 1800);
  } finally {
    loading.value = false;
  }
};

onMounted(async () => {
  const { username } = route.query;
  if (username) {
    await handleQrCode2Fa(username as string);
    user.value = username as string;
  }
});
</script>

<template>
  <v-container class="flex items-center justify-center min-h-screen">
    <v-card class="mx-auto w-full md:w-1/2 lg:w-1/2 xl:w-[40%]">
      <v-card-title>
        <h1 class="text-h5 font-semibold text-primary pt-6 text-center">
          QR Code
        </h1>
        <h1 class="text-md font-semibold text-center text-neutral">
          Google Authentication two factor
        </h1>
        <div v-if="qrCodeUrl" class="flex justify-center w-full mt-4 mb-6">
          <qrcode-vue :size="150" :value="qrCodeUrl" />
        </div>
        <v-text-field
          class="mt-7"
          v-model="user"
          variant="outlined"
          label="ชื่อผู้ใช้งาน"
          prepend-inner-icon="mdi-account"
        ></v-text-field>

        <div class="w-full">
          <v-btn
            :disabled="user === ''"
            block
            :loading="loading"
            @click="handleQrCode2Fa(user)"
            >ขอ Qr Code 2Fa
          </v-btn>
        </div>

        <v-btn
          variant="outlined"
          class="mt-3"
          block
          :loading="loading"
          to="/auth/signIn"
        >
          เข้าสู่ระบบ
        </v-btn>
      </v-card-title>
      <v-card-text>
        <p class="text-[13px] mt-3 indent-8">
          เนื่องจากความสำคัญของการรักษาความปลอดภัยให้กับระบบ การใช้
          <span class="text-primary"
            >Google Authenticator two-factor authentication (2FA)</span
          >
          จึงเป็นตัวเลือกที่ดีในการป้องกันการเข้าถึงที่ไม่ได้รับอนุญาต
          ขั้นตอนสำหรับขอการใช้
          <span class="text-primary">QR code 2FA</span> มีดังนี้
        </p>
        <ol class="list-decimal list-inside text-[13px]">
          <li class="text-[13px]">
            เปิดแอป
            <span class="text-primary">Google Authenticator</span>
            บนอุปกรณ์ของคุณ และสแกน
            <span class="text-primary">QR code</span> ที่ปรากฏ
          </li>
          <li class="text-[13px]">
            ป้อนรหัสยืนยันตัวตนที่แสดงในแอป
            <span class="text-primary">Google Authenticator</span>
          </li>
          <li class="text-[13px]">
            ไปที่หน้าเข้าสู่ระบบเเล้วกรอกชื่อผู้ใช้งานงานเเละรหัสผ่านให้ครบ
            จากนั้นคลิก <br />
            "เข้าสู่ระบบ"
          </li>
          <li class="text-[13px]">
            นำหมายเลขที่ปรากฏใน
            <span class="text-primary">Google Authenticator</span>
            ไปกรอกเพื่อยืนยันรหัสผ่าน (2FA) ในป๊อปอั๊ป
          </li>
        </ol>
      </v-card-text>
    </v-card>
  </v-container>
</template>
