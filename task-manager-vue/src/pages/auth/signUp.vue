<script lang="ts" setup>
import { useAuthStore } from "@/stores/auth";
import { AuthSignUp } from "@/types/Auth";
import { toast } from "vue3-toastify";

const formRef: any = ref(null);
const initForm = reactive<AuthSignUp>({
  full_name: "",
  username: "",
  password: "",
  confirmPassword: "",
});
const passwordRef = ref<any>(null);
const confirmPasswordRef = ref<any>(null);
const authStore = useAuthStore();
let formData = reactive(initForm);

const rules = {
  full_name: [
    (value: string) => {
      if (value) return true;
      return "กรุณากรอกชื่อ-นามสกุล";
    },
  ],
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
  confirmPassword: [
    (value: string) => {
      if (value) return true;
      return "กรุณากรอกยืนยันรหัสผ่าน";
    },
  ],
};

const openPassword = (key = "password") => {
  let input = passwordRef.value.$el.querySelector("input");
  if (key === "confirmPassword") {
    input = confirmPasswordRef.value.$el.querySelector("input");
  }
  input.type = input.type === "password" ? "text" : "password";
};

const handleSubmit = async () => {
  authStore.auth.loading = true;
  try {
    const { valid } = await formRef.value.validate();
    if (!valid) {
      toast.error("กรุณากรอกข้อมูลให้ครบ");
      return false;
    }

    if (formData.confirmPassword !== formData.password) {
      toast.error("รหัสผ่านไม่ตรงกัน");
      return false;
    }

    const payload: AuthSignUp = {
      ...formData,
    };
    await authStore.handleSignUp(payload);
  } catch (error) {
    console.error(error);
  } finally {
    authStore.auth.loading = false;
  }
};
</script>

<template>
  <v-container class="flex items-center justify-center min-h-screen">
    <v-card class="mx-auto w-full md:w-1/2 lg:w-1/2 xl:w-1/3">
      <v-form ref="formRef" @submit.prevent="handleSubmit">
        <v-card-title>
          <h1 class="text-h6 font-semibold text-primary mt-2 text-center">
            Task Management System
          </h1>
          <h1 class="text-md my-2 font-semibold text-center text-neutral">
            ลงทะเบียนผู้ใช้
          </h1>
        </v-card-title>
        <v-card-text>
          <section class="mt-1 flex flex-col w-full h-full gap-2">
            <v-text-field
              clearable
              v-model="formData.full_name"
              :rules="rules.full_name"
              variant="outlined"
              label="ชื่อ-นามสกุล"
              prepend-inner-icon="mdi-account"
            />
            <v-text-field
              clearable
              v-model="formData.username"
              :rules="rules.username"
              variant="outlined"
              label="ชื่อผู้ใช้"
              prepend-inner-icon="mdi-account"
            />
            <v-text-field
              id="password"
              v-model="formData.password"
              :rules="rules.password"
              ref="passwordRef"
              type="password"
              variant="outlined"
              label="รหัสผ่าน"
              prepend-inner-icon="mdi-lock"
            >
              <template #append-inner>
                <v-icon icon="mdi-eye" @click="openPassword"></v-icon>
              </template>
            </v-text-field>
            <v-text-field
              id="confirmPassword"
              v-model="formData.confirmPassword"
              :rules="rules.confirmPassword"
              type="password"
              ref="confirmPasswordRef"
              variant="outlined"
              label="ยืนยันรหัสผ่าน"
              prepend-inner-icon="mdi-lock"
            >
              <template #append-inner>
                <v-icon
                  icon="mdi-eye"
                  @click="openPassword('confirmPassword')"
                ></v-icon>
              </template>
            </v-text-field>
          </section>

          <section class="text-center mt-2">
            <v-btn
              block
              color="primary"
              type="submit"
              :loading="authStore.auth.loading"
              >ลงทะเบียน
            </v-btn>
            <p class="text-neutral mt-5">หรือ</p>
            <v-btn block variant="outlined">
              <v-icon icon="mdi-google"></v-icon>
              <span class="ml-2 font-semibold">ลงทะเบียนด้วย Google</span>
            </v-btn>
            <v-btn block variant="outlined" class="mt-3">
              <v-icon icon="mdi-github"></v-icon>
              <span class="ml-2 font-semibold">ลงทะเบียนด้วย Github</span>
            </v-btn>
            <p class="mt-6 font-semibold text-neutral text-sm">
              ถ้ามีบัญชีผู้ใช้แล้ว
              <RouterLink to="/auth/signIn">
                <span class="text-primary cursor-pointer"
                  >เข้าสู่ระบบที่นี่
                </span>
              </RouterLink>
            </p>
          </section>
        </v-card-text>
      </v-form>
    </v-card>
  </v-container>
</template>
