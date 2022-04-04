import {useAxios} from "@vueuse/integrations/useAxios";
import baseAxios from "@/apis/base";

export default function (id: Number) {
    // @todo fix data params for ts
    // @ts-ignore
    return useAxios(`/tasks/${id}/`, baseAxios)
}