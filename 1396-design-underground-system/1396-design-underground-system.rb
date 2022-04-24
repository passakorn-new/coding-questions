class UndergroundSystem
    def initialize()
        @check = {}
        @time = {}
    end

    def check_in(id, station_name, t)
        @check[id] = [station_name, t]
    end

    def check_out(id, station_name, t)
        station_prev, t_prev = @check[id]
        t_diff = t - t_prev
        
        if @time[[station_prev, station_name]]
          total_time, station_i = @time[[station_prev, station_name]]
          @time[[station_prev, station_name]] = [total_time + t_diff, station_i + 1]
        else
          @time[[station_prev, station_name]] = [t_diff, 1]
        end
    end

    def get_average_time(start_station, end_station)
        total_time, station_i = @time[[start_station, end_station]]
        total_time.to_f/station_i
    end
end

# Your UndergroundSystem object will be instantiated and called as such:
# obj = UndergroundSystem.new()
# obj.check_in(id, station_name, t)
# obj.check_out(id, station_name, t)
# param_3 = obj.get_average_time(start_station, end_station)