<?php

use Illuminate\Database\Migrations\Migration;
use Illuminate\Database\Schema\Blueprint;
use Illuminate\Support\Facades\Schema;

class CreateHopdonglaodongTable extends Migration
{
    /**
     * Run the migrations.
     *
     * @return void
     */
    public function up()
    {
        Schema::create('hopdonglaodong', function (Blueprint $table) {
            $table->id();
            $table->string('name');
            $table->date('ngaybatdau');
            $table->date('thoigianlaodong');
            $table->text('ghichu');
            $table->timestamps();
        });
    }

    /**
     * Reverse the migrations.
     *
     * @return void
     */
    public function down()
    {
        Schema::dropIfExists('hopdonglaodong');
    }
}
