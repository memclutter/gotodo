export default async ({store, app: {$services}}) => {
  await store.dispatch('init', $services);
}
