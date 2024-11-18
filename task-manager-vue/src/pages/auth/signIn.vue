<script lang="ts" setup>
import { useAuthStore } from "@/stores/auth";
import { AuthSignIn } from "@/types/Auth";
import { toast } from "vue3-toastify";

const formRef: any = ref(null);
const passwordRef: any = ref(null);
const dialog2Fa = ref(false);
const authStore = useAuthStore();
const initForm = reactive<AuthSignIn>({
  username: "",
  password: "",
});

let formData = reactive(initForm);

const rules = {
  username: [
    (value: string) => {
      if (value) return true;
      return "กรุณากรอกชื่อผู้ใช้งาน";
    },
  ],
  password: [
    (value: string) => {
      if (value) return true;
      return "กรุณากรอกรหัสผ่าน";
    },
  ],
};

const openPassword = () => {
  const input = passwordRef.value.$el.querySelector("input");
  input.type = input.type === "password" ? "text" : "password";
};

const handleSubmit = async () => {
  try {
    const { valid } = await formRef.value.validate();
    if (!valid) {
      toast.error("กรุณากรอกข้อมูลให้ครบ");
      return false;
    }
    dialog2Fa.value = true;
  } catch (error) {
    console.error(error);
  }
};

const handleSignIn2fa = async (otp: string) => {
  const payload = { ...formData, otp_code: otp };
  await authStore.handleSignIn(payload);
};
</script>

<template>
  <v-container class="flex items-center justify-center min-h-screen">
    <v-card class="mx-auto w-full md:w-1/2 lg:w-1/2 xl:w-1/3">
      <v-form ref="formRef" @submit.prevent="handleSubmit">
        <v-card-title>
          <h1 class="text-h6 font-semibold text-primary py-3 mt-2 text-center">
            Task Management System
          </h1>
          <h1 class="text-md my-2 font-semibold text-center text-neutral">
            เข้าสู่ระบบ
          </h1>
        </v-card-title>
        <v-card-text>
          <section class="mt-1 flex flex-col w-full h-full gap-2">
            <v-text-field
              clearable
              v-model="formData.username"
              :rules="rules.username"
              variant="outlined"
              label="ชื่อผู้ใช้"
              prepend-inner-icon="mdi-account"
            ></v-text-field>
            <v-text-field
              ref="passwordRef"
              v-model="formData.password"
              :rules="rules.password"
              type="password"
              variant="outlined"
              label="รหัสผ่าน"
              prepend-inner-icon="mdi-lock"
            >
              <template #append-inner>
                <v-icon icon="mdi-eye" @click="openPassword"></v-icon>
              </template>
            </v-text-field>
          </section>

          <section class="text-center mt-3">
            <v-btn block color="primary" type="submit">เข้าสู่ระบบ</v-btn>
            <p class="text-primary my-4 cursor-pointer">ลืมรหัสผ่าน?</p>
            <p class="text-neutral">หรือ</p>
            <v-btn block variant="outlined">
              <v-icon icon="mdi-google"></v-icon>
              <span class="ml-2 font-semibold">เข้าสู่ระบบด้วย Google</span>
            </v-btn>
            <v-btn block variant="outlined" class="mt-3">
              <v-icon icon="mdi-github"></v-icon>
              <span class="ml-2 font-semibold">เข้าสู่ระบบด้วย Github</span>
            </v-btn>
            <p class="mt-6 font-semibold text-neutral text-sm">
              หากยังไม่มีบัญชีกรุณา
              <RouterLink to="/auth/signUp">
                <span class="text-primary cursor-pointer">ลงทะเบียนที่นี่</span>
              </RouterLink>
            </p>
          </section>
        </v-card-text>
      </v-form>
    </v-card>
  </v-container>

  <Auth2FADialog
    :username="formData.username"
    v-model:isDialogVisible="dialog2Fa"
    @confirm="handleSignIn2fa"
  />
</template>
