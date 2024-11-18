export const vuetifyConfig = {
  theme: {
    defaultTheme: "light",
    themes: {
      light: {
        dark: false,
        colors: {
          primary: "#6366F1", // Indigo
          secondary: "#EC4899", // Pink
          accent: "#8B5CF6", // Purple
          success: "#10B981", // Emerald
          warning: "#F59E0B", // Amber
          error: "#EF4444", // Red
          info: "#3B82F6", // Blue
          background: "#F9FAFB",
          surface: "#FFFFFF",
          "primary-darken-1": "#4F46E5",
          "surface-variant": "#F3F4F6",
          "on-surface-variant": "#374151",
          neutral: "#4B5563",
        },
        variables: {
          // "border-color": "#E5E7EB",
          // "border-opacity": 0.12,
          "high-emphasis-opacity": 0.87,
          "medium-emphasis-opacity": 0.6,
          "disabled-opacity": 0.38,
          "idle-opacity": 0.04,
          "hover-opacity": 0.04,
          "focus-opacity": 0.12,
          "selected-opacity": 0.08,
          "activated-opacity": 0.12,
          "pressed-opacity": 0.12,
          "dragged-opacity": 0.08,
          "theme-kbd": "#212529",
          "theme-on-kbd": "#FFFFFF",
          "theme-code": "#F8F9FA",
          "theme-on-code": "#212529",
        },
      },
      dark: {
        dark: true,
        colors: {
          primary: "#818CF8", // Indigo lighter
          secondary: "#F472B6", // Pink lighter
          accent: "#A78BFA", // Purple lighter
          success: "#34D399", // Emerald lighter
          warning: "#FBBF24", // Amber lighter
          error: "#F87171", // Red lighter
          info: "#60A5FA", // Blue lighter
          background: "#111827",
          surface: "#1F2937",
          "primary-darken-1": "#6366F1",
          "surface-variant": "#374151",
          "on-surface-variant": "#D1D5DB",
          neutral: "#D1D5DB",
        },
        variables: {
          // "border-color": "#374151",
          // "border-opacity": 0.12,
          "high-emphasis-opacity": 1,
          "medium-emphasis-opacity": 0.7,
          "disabled-opacity": 0.38,
          "idle-opacity": 0.1,
          "hover-opacity": 0.1,
          "focus-opacity": 0.12,
          "selected-opacity": 0.12,
          "activated-opacity": 0.12,
          "pressed-opacity": 0.16,
          "dragged-opacity": 0.12,
          "theme-kbd": "#212529",
          "theme-on-kbd": "#FFFFFF",
          "theme-code": "#343A40",
          "theme-on-code": "#F8F9FA",
        },
      },
    },
  },
};
