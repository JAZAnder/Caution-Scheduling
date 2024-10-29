function getUserHourByAnyFilter(hourId, tutorId, dayOfWeek) {
    const { data, loading, error } = useFetch(
        `/api/availability?tutorId=`+tutorId+`&hourId=`+hourId+`&dayOfWeek=`+dayOfWeek,
        { method: "get" },
        []
      );


    return(data, loading, error)
}