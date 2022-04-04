import {useAxios} from "@vueuse/integrations/useAxios";
import baseAxios from "@/apis/base";

export default function () {
    // @todo fix data params for ts
    // @ts-ignore
    return useAxios('/tasks/', baseAxios)
}