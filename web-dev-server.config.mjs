export default {
  open: true,
  watch: true,
  rootDir: 'src',
  nodeResolve: {
    exportConditions: ['development'],
    dedupe: true,
  },
  esbuildTarget: 'auto',
};
