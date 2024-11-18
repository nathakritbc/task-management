<script setup lang="ts">
import { useAuthStore } from "@/stores/auth";
import { defineEmits, defineProps } from "vue";
import { toast } from "vue3-toastify";

const props = defineProps({
  username: {
    type: String,
    required: true,
  },
  isDialogVisible: {
    default: false,
    type: Boolean,
    required: true,
  },
});

const emit = defineEmits(["update:isDialogVisible", "confirm"]);
const authStore = useAuthStore();
const otp = ref("");
const errorForm = ref(false);
const dialog = computed({
  get() {
    return props.isDialogVisible;
  },
  set(val: boolean) {
    emit("update:isDialogVisible", val);
  },
});

const handleSubmit = async () => {
  errorForm.value = false;
  try {
    if (otp.value.length < 6) {
      toast.error("กรุณากรอกรหัสผ่านให้ครบ 6 หลัก");
      errorForm.value = true;
      return false;
    }
    emit("confirm", otp.value);
  } catch (error) {
    console.error(error);
  }
};

watch(
  () => props.isDialogVisible,
  (newValue: boolean) => {
    console.log("newValue", newValue);
    otp.value = "";
  }
);
</script>

<template>
  <v-dialog max-width="500" v-model="dialog">
    <v-card
      class="py-8 px-6 text-center mx-auto ma-4"
      elevation="12"
      max-width="400"
      width="100%"
    >
      <h3 class="text-h6 mb-3 text-primary font-semibold">
        ยืนยันรหัสผ่าน (2FA)
      </h3>

      <div class="text-body-2 mb-1">
        กรุณากรอกรหัสผ่านยืนยันตัวตน 6 หลัก <br />ที่ได้รับจาก
        <span class="text-primary font-semibold">Google Authenticator</span>
        เพื่อเข้าสู่ระบบ
      </div>

      <v-form @submit.prevent="handleSubmit">
        <v-sheet color="surface">
          <v-otp-input
            autofocus
            v-model="otp"
            type="number"
            variant="outlined"
          ></v-otp-input>
          <p class="text-error" v-if="errorForm">กรุณากรอกรหัสผ่าน</p>
        </v-sheet>

        <v-btn
          block
          :loading="authStore.auth.loading"
          type="submit"
          class="mt-2"
          color="primary"
          height="40"
          text="ยืนยัน"
          variant="flat"
        ></v-btn>
        <v-btn
          block
          class="mt-2 mb-5"
          :to="'/auth/regis2Fa?username=' + username"
          color="primary"
          height="40"
          variant="outlined"
          >ขอรับรหัสผ่าน (2FA)</v-btn
        >
      </v-form>
    </v-card>
  </v-dialog>
</template>
